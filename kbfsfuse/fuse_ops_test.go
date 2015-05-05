package main

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/kbfs/libkbfs"
	"testing"
)

// Given the list of users, create and return a config suitable for
// unit-testing.
func makeTestConfig(users []string) *libkbfs.ConfigLocal {
	config := libkbfs.NewConfigLocal()

	localUsers := make([]libkbfs.LocalUser, len(users))
	for i := 0; i < len(users); i++ {
		kid := libkbfs.KID("test_sub_key_" + users[i])
		localUsers[i] = libkbfs.LocalUser{
			Name:            users[i],
			Uid:             libkb.UID{byte(i + 1)},
			SubKeys:         []libkbfs.Key{libkbfs.NewKeyFake(kid)},
			DeviceSubkeyKid: kid,
		}
	}
	loggedInUid := localUsers[0].Uid

	// TODO: Consider using fake BlockOps and MDOps instead.
	k := libkbfs.NewKBPKILocal(loggedInUid, localUsers)
	config.SetKBPKI(k)
	config.SetBlockServer(libkbfs.NewFakeBlockServer())
	config.SetMDServer(libkbfs.NewFakeMDServer(config))

	return config
}

func doLookupOrBust(t *testing.T, parent *FuseNode, name string) *FuseNode {
	var attr fuse.Attr
	inode, code := parent.Lookup(&attr, name, nil)
	if code != fuse.OK {
		t.Fatalf("Lookup failure: %s", code)
	}
	return inode.Node().(*FuseNode)
}

func doMkDirOrBust(t *testing.T, parent *FuseNode, name string) *FuseNode {
	inode, code := parent.Mkdir(name, 0, nil)
	if code != fuse.OK {
		t.Fatalf("Mkdir failure: %s", code)
	}
	return inode.Node().(*FuseNode)
}

// Test that looking up one's own public directory works.
func TestLookupSelfPublic(t *testing.T) {
	config := makeTestConfig([]string{"test_user"})

	root := NewFuseRoot(config)
	defer root.Ops.Shutdown()
	_ = nodefs.NewFileSystemConnector(root, nil)

	node1 := doLookupOrBust(t, root, "test_user")
	doLookupOrBust(t, node1, "public")
}

// Test that looking up someone else's public directory works.
func TestLookupOtherPublic(t *testing.T) {
	config := makeTestConfig([]string{"test_user1", "test_user2"})

	// First, look up the test_user1/public as test_user1 to
	// create it.

	root := NewFuseRoot(config)
	defer root.Ops.Shutdown()
	_ = nodefs.NewFileSystemConnector(root, nil)

	node1 := doLookupOrBust(t, root, "test_user1")
	doLookupOrBust(t, node1, "public")

	// Now, simulate a remount as test_user2.

	config.KBPKI().(*libkbfs.KBPKILocal).LoggedIn = libkb.UID{2}
	config.SetMDCache(libkbfs.NewMDCacheStandard(5000))
	root = NewFuseRoot(config)
	_ = nodefs.NewFileSystemConnector(root, nil)

	// Then, do the lookup again as test_user2.

	node1 = doLookupOrBust(t, root, "test_user1")
	doLookupOrBust(t, node1, "public")
}

// Test that nodes start off as not needing updating, making a
// directory marks the path from the new directory's parent to the
// user root (in this case just one directory) as needing updating,
// and looking up a directory that needs updating marks it as not
// needing updating.
func TestNeedUpdateBasic(t *testing.T) {
	config := makeTestConfig([]string{"test_user"})

	root := NewFuseRoot(config)
	defer root.Ops.Shutdown()
	_ = nodefs.NewFileSystemConnector(root, nil)

	if root.NeedUpdate {
		t.Error("/ unexpectedly needs update")
	}

	// Look up /test_user.
	node1 := doLookupOrBust(t, root, "test_user")

	if root.NeedUpdate {
		t.Error("/ unexpectedly needs update")
	}

	if node1.NeedUpdate {
		t.Error("/test_user unexpectedly needs update")
	}

	// Make /test_user/dir.
	node2 := doMkDirOrBust(t, node1, "dir")

	if root.NeedUpdate {
		t.Error("/ unexpectedly needs update")
	}

	if !node1.NeedUpdate {
		t.Error("/test_user unexpectedly doesn't need update")
	}

	if node2.NeedUpdate {
		t.Error("/test_user/dir unexpectedly needs update")
	}

	// Look up /test_user again.
	node1 = doLookupOrBust(t, root, "test_user")

	if root.NeedUpdate {
		t.Error("/ unexpectedly needs update")
	}

	if node1.NeedUpdate {
		t.Error("/test_user unexpectedly needs update")
	}

	if node2.NeedUpdate {
		t.Error("/test_user/dir unexpectedly needs update")
	}
}

// Test that nodes start off as not needing updating, making a
// directory marks the path from the new directory's parent to the
// user root as needing updating, and not just the parent.
func TestNeedUpdateAll(t *testing.T) {
	config := makeTestConfig([]string{"test_user"})

	root := NewFuseRoot(config)
	defer root.Ops.Shutdown()
	_ = nodefs.NewFileSystemConnector(root, nil)

	// Make /test_user/dir1/dir2/dir3 and clear their NeedUpdate
	// flags.

	node1 := doLookupOrBust(t, root, "test_user")
	node2 := doMkDirOrBust(t, node1, "dir1")
	node3 := doMkDirOrBust(t, node2, "dir2")
	node4 := doMkDirOrBust(t, node3, "dir3")

	node1 = doLookupOrBust(t, root, "test_user")
	node2 = doLookupOrBust(t, node1, "dir1")
	node3 = doLookupOrBust(t, node2, "dir2")
	node4 = doLookupOrBust(t, node3, "dir3")

	// Make /test_user/dir1/dir2/dir3/dir4.
	node5 := doMkDirOrBust(t, node4, "dir4")

	if root.NeedUpdate {
		t.Error("/ unexpectedly needs update")
	}

	if !node1.NeedUpdate {
		t.Error("/test_user unexpectedly doesn't need update")
	}

	if !node2.NeedUpdate {
		t.Error("/test_user/dir1 unexpectedly doesn't need update")
	}

	if !node3.NeedUpdate {
		t.Error("/test_user/dir2 unexpectedly doesn't need update")
	}

	if !node4.NeedUpdate {
		t.Error("/test_user/dir3 unexpectedly doesn't need update")
	}

	if node5.NeedUpdate {
		t.Error("/test_user/dir4 unexpectedly needs update")
	}
}
