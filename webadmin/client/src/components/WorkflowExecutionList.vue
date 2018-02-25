<template>
  <div>
    <h3>WorkflowExecution List</h3>
    <div>
      <table class="highlight bordered">
        <thead>
        <tr>
          <th>ID</th>
          <th>Status</th>
          <th>Workflow Name</th>
          <th>Execution Name</th>
          <th>Start</th>
          <th>End</th>
        </tr>
        </thead>

        <tbody>
        <tr v-for="we in workflowExecs" :key="we.id" v-on:click="onClickExec(we.id)">
          <td>{{ we.id }}</td>
          <td>{{ we.status }}</td>
          <td>{{ we.name }}</td>
          <td>{{ we.workflow.name }}</td>
          <td>{{ we.startedAt }}</td>
          <td>{{ we.endedAt }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import JobnetesApi from '@/external/jobnetesApi'
import Promise from 'es6-promise'

export default {
  name: 'WorkflowExecutionList',
  data () {
    return {
      workflowExecs: []
    }
  },
  beforeRouteEnter (route, redirect, next) {
    JobnetesApi.getWorkflowExecutions()
      .then(response => response.json())
      .then(data => {
        next(vm => {
          console.log(Promise.resolve())
          vm.workflowExecs = data.items
        })
      })
  },
  beforeRouteUpdate (to, from, next) {
    this.workflowExecs = []
    JobnetesApi.getWorkflowExecutions()
      .then(response => response.json())
      .then(data => {
        this.workflowExecs = data.items
        next()
      })
  },
  methods: {
    onClickExec: function (wid) {
      this.$router.push({
        name: 'WorkflowExecutionDetail',
        params: {
          id: wid
        }
      })
    }
  }
}
</script>
