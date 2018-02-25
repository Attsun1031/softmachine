'use strict'
import 'whatwg-fetch'

export default {
  getWorkflowExecutions: () => {
    return fetch('/api/v1/workflow/execution')
  },
  getWorkflowExecutionDetail: (wid) => {
    return fetch('/api/v1/workflow/execution/' + encodeURIComponent(wid))
  }
}
