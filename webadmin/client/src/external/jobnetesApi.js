'use strict'
import 'whatwg-fetch'

export default {
  getWorkflowExecutions: () => {
    return fetch('/api/v1/workflow/execution')
  }
}
