<style scoped>
  #mynetwork {
    width: 600px;
    height: 400px;
    border: 1px solid lightgray;
  }
</style>
<template>
  <div>
    <h3>WorkflowExecution Detail</h3>
    <div v-if="wfExec !== null">
      <div>
        <div>
          <h4>Workflow</h4>
          <table class="highlight bordered">
            <thead>
            <tr>
              <th>Status</th>
              <th>Workflow Name</th>
              <th>Execution Name</th>
              <th>Start</th>
              <th>End</th>
            </tr>
            </thead>

            <tbody>
            <tr>
              <td>{{ wfExec.status }}</td>
              <td>{{ wfExec.name }}</td>
              <td>{{ wfExec.workflow.name }}</td>
              <td>{{ wfExec.startedAt }}</td>
              <td>{{ wfExec.endedAt }}</td>
            </tr>
            </tbody>
          </table>
        </div>
        <div>
          <h4>Task</h4>
          <table class="highlight bordered">
            <thead>
            <tr>
              <th>ID</th>
              <th>Status</th>
              <th>Task Name</th>
              <th>Execution Name</th>
              <th>Start</th>
              <th>End</th>
            </tr>
            </thead>

            <tbody>
            <tr v-for="task in wfExec.taskExecutions" :key="task.id">
              <td>{{ task.id }}</td>
              <td>{{ task.status }}</td>
              <td>{{ task.taskName }}</td>
              <td>{{ task.executionName }}</td>
              <td>{{ task.startedAt }}</td>
              <td>{{ task.endedAt }}</td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div>
        <h4>Dependency Graph</h4>
        <div id="mynetwork" ref="mynetwork"></div>
      </div>
    </div>
  </div>
</template>

<script>
import JobnetesApi from '@/external/jobnetesApi'
import vis from 'vis'

export default {
  name: 'WorkflowExecutionDetail',
  props: ['id'],
  data: function () {
    return {
      wfExec: null,
      network: null
    }
  },
  beforeRouteEnter: function (route, redirect, next) {
    next(vm => {
      JobnetesApi.getWorkflowExecutionDetail(vm.id)
        .then(response => response.json())
        .then(data => {
          vm.wfExec = data.item
        })
    })
  },
  beforeRouteUpdate: function (to, from, next) {
    this.wfExec = null
    JobnetesApi.getWorkflowExecutionDetail(this.id)
      .then(response => response.json())
      .then(data => {
        this.wfExec = data.item
        next()
      })
  },
  updated: function () {
    // TODO: vis.jsのサンプル
    let container = this.$refs.mynetwork
    if (this.network !== null || container === undefined || container === null) {
      return
    }

    let nodes = new vis.DataSet([
      {id: 1, label: 'Node 1'},
      {id: 2, label: 'Node 2'},
      {id: 3, label: 'Node 3'},
      {id: 4, label: 'Node 4'},
      {id: 5, label: 'Node 5'}
    ])

    // create an array with edges
    let edges = new vis.DataSet([
      {from: 1, to: 3},
      {from: 1, to: 2},
      {from: 2, to: 4},
      {from: 2, to: 5}
    ])

    // provide the data in the vis format
    let d = {
      nodes: nodes,
      edges: edges
    }
    let options = {
    }

    // initialize your network!
    this.network = new vis.Network(container, d, options)
    this.network.on('click', function (properties) {
      let ids = properties.nodes
      let clickedNodes = nodes.get(ids)
      console.log('clicked nodes:', clickedNodes)
    })
  }
}
</script>
