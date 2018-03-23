'use strict'
import 'whatwg-fetch'

export default {
  getWorkflowExecutions: () => {
    return fetch('/api/v1/workflow/execution')
  },
  getWorkflowExecutionDetail: (wid) => {
    return fetch('/api/v1/workflow/execution/' + encodeURIComponent(wid))
  },
  getPods: (tid) => {
    return fetch('/api/v1/task/' + encodeURIComponent(tid) + '/pod')
  },
  getContainerLog: (tid, pod, container) => {
    return fetch('/api/v1/task/' +
      encodeURIComponent(tid) +
      '/pod/' +
      encodeURIComponent(pod) +
      '/' +
      encodeURIComponent(container) +
      '/log')
  }
}
