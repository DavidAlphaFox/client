// @flow

import React, {Component} from 'react'
import {globalMargins, globalColors} from '../../styles'
import * as Kb from '../../common-adapters'
import {getOtherErrorInfo} from './index.js'

type Props = {
  onSendFeedbackContained: () => void,
  showSuccessBanner: boolean,
  sendLogs: boolean,
  feedback: ?string,
  sending: boolean,
  sendError: ?Error,
  onChangeSendLogs: (nextValue: boolean) => void,
  onChangeFeedback: (nextValue: ?string) => void,
}

class Feedback extends Component<Props> {
  _onSubmit = () => {
    this.props.onSendFeedbackContained()
  }

  render() {
    const {
      showSuccessBanner,
      sendLogs,
      onChangeSendLogs,
      feedback,
      onChangeFeedback,
      sending,
      sendError,
    } = this.props
    return (
      <Kb.Box2 direction="vertical" fullWidth={true} alignItems="center">
        {showSuccessBanner && (
          <Kb.Box2
            alignItems="center"
            direction="horizontal"
            fullWidth={true}
            style={{
              backgroundColor: globalColors.green,
              flex: 0,
              minHeight: 40,
              paddingBottom: globalMargins.tiny,
              paddingTop: globalMargins.tiny,
            }}
          >
            <Kb.Text center={true} type="BodySmallSemibold" negative={true} style={{flex: 1}}>
              Thanks! Your feedback was sent.
            </Kb.Text>
          </Kb.Box2>
        )}
        <Kb.Box2
          direction="vertical"
          style={{
            alignItems: 'stretch',
            justifyContent: 'flex-start',
            marginLeft: globalMargins.small,
            marginRight: globalMargins.small,
          }}
        >
          <Kb.Box2
            direction="horizontal"
            fullWidth={true}
            style={{
              paddingTop: globalMargins.small,
            }}
          >
            <Kb.Input
              autoCapitalize="sentences"
              autoCorrect={true}
              autoFocus={true}
              style={{flex: 1}}
              hideLabel={true}
              inputStyle={{textAlign: 'left'}}
              multiline={true}
              rowsMin={3}
              rowsMax={10}
              hintText="Please tell us what you were doing, your experience, or anything else we should know. Thanks!"
              value={feedback}
              onChangeText={onChangeFeedback}
            />
          </Kb.Box2>
          <Kb.Box2 direction="horizontal">
            <Kb.Checkbox
              label=""
              style={{alignItems: 'flex-start', marginRight: globalMargins.tiny}}
              checked={sendLogs}
              onCheck={onChangeSendLogs}
            />
            <Kb.Box2 direction="vertical">
              <Kb.Text type="Body">Include your logs</Kb.Text>
              <Kb.Text type="BodySmall">
                This includes some private metadata info (e.g., filenames, but not contents) but it will help
                the developers fix bugs more quickly.
              </Kb.Text>
            </Kb.Box2>
          </Kb.Box2>
          <Kb.ButtonBar style={{paddingTop: globalMargins.small}}>
            <Kb.Button fullWidth={true} label="Send" onClick={this._onSubmit} waiting={sending} />
          </Kb.ButtonBar>
          {sendError && (
            <Kb.Box2 direction="vertical" style={{marginTop: globalMargins.small}}>
              <Kb.Text type="BodySmallError">Could not send log</Kb.Text>
              <Kb.Text type="BodySmall" selectable={true} style={{marginBottom: 10, marginTop: 10}}>{`${
                sendError.name
              }: ${sendError.message}`}</Kb.Text>
              <Kb.Text type="BodySmallSemibold">Stack</Kb.Text>
              <Kb.Text type="BodySmall" selectable={true} style={{marginBottom: 10, marginTop: 10}}>
                {sendError.stack}
              </Kb.Text>
              <Kb.Text type="BodySmallSemibold">Error dump</Kb.Text>
              <Kb.Text type="BodySmall" selectable={true} style={{marginBottom: 10, marginTop: 10}}>
                {JSON.stringify(getOtherErrorInfo(sendError), null, 2)}
              </Kb.Text>
            </Kb.Box2>
          )}
        </Kb.Box2>
      </Kb.Box2>
    )
  }
}

export default Feedback
