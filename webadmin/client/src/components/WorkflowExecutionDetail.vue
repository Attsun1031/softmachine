<template>
  <div>
    <h3>WorkflowExecution Detail</h3>
    <div v-if="wfExec !== null">
      <div>
        <h4>Task Execution Graph</h4>
        <svg id="workflowDag">
          <g></g>
        </svg>
      </div>
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
            <tr v-for="task in wfExec.taskExecutions" :key="task.id" v-on:click="onClickTask(task.id)">
              <td>{{ task.id }}</td>
              <td>{{ getStatusLabel(task.status) }}</td>
              <td>{{ task.taskName }}</td>
              <td>{{ task.executionName }}</td>
              <td>{{ task.startedAt }}</td>
              <td>{{ task.endedAt }}</td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import JobnetesApi from '@/external/jobnetesApi'
import dagreD3 from 'dagreD3'
import * as d3 from 'd3'

export default {
  name: 'WorkflowExecutionDetail',
  props: ['id'],
  data: function () {
    return {
      wfExec: null,
      dag: null
    }
  },
  methods: {
    onClickTask: function (tid) {
      this.$router.push({
        name: 'TaskExecutionLog',
        params: {
          taskId: tid
        }
      })
    },
    isParallelJob: function (te) {
      return te.taskType === 'parallel-job'
    },
    getParallelJobEndNodeId: function (te) {
      return -te
    },
    getStatusLabel: function (status) {
      switch (status) {
        case 0:
          return 'running'
        case 1:
          return 'success'
        case 2:
          return 'failure'
        default:
          return 'unknown'
      }
    },
    getShape: function (te) {
      if (this.isParallelJob(te)) {
        return 'ellipse'
      } else {
        return 'rect'
      }
    },
    renderDag: function () {
      if (this.dag !== null || this.wfExec === null) {
        return
      }
      let g = new dagreD3.graphlib.Graph()
      this.dag = g
      g.setGraph({
        rankdir: 'LR'
      })
      g.setDefaultEdgeLabel(function () {
        return {}
      })

      // collect parents
      let parentJobs = new Map()
      this.wfExec.taskExecutions.filter(te => te.parentId > 0).forEach(te => {
        if (!parentJobs.has(te.parentId)) {
          parentJobs.set(te.parentId, [])
        }
        parentJobs.get(te.parentId).push(te)
      })

      // TODO: 並列タスクの完了を示すノードを作る
      this.wfExec.taskExecutions.forEach(te => {
        g.setNode(te.id, {label: te.taskName, class: this.getStatusLabel(te.status), shape: this.getShape(te)})
        if (parentJobs.has(te.id) && this.getStatusLabel(te.status) !== 'running') {
          g.setNode(this.getParallelJobEndNodeId(te), {label: 'END OF PARALLEL', class: this.getStatusLabel(te.status), shape: this.getShape(te)})
        }

        // draw parent node to initial child nodes
        if (te.parentId > 0 && te.prevId === 0) {
          g.setEdge(te.parentId, te.id)
        }

        // draw to next
        if (te.nextId > 0) {
          if (parentJobs.has(te.id)) {
            // draw end node to next
            let endNodeId = this.getParallelJobEndNodeId(te)
            g.setEdge(endNodeId, te.nextId)
            // draw next from end children
            parentJobs.get(te.id).filter(te => te.nextId === 0).forEach(ec => {
              g.setEdge(ec.id, endNodeId)
            })
          } else {
            g.setEdge(te.id, te.nextId)
          }
        }
      })

      g.nodes().forEach(function (v) {
        if (v !== undefined) {
          let node = g.node(v)
          // Round the corners of the nodes
          node.rx = node.ry = 5
        }
      })

      // eslint-disable-next-line
      let render = new dagreD3.render()

      let svg = d3.select('#workflowDag')
      render(d3.select('#workflowDag g'), g)
      svg.attr('height', g.graph().height)
      svg.attr('width', g.graph().width)

      d3.selectAll('svg rect')
        .on('mouseover', function (d) {
          console.log('mouseover')
        })
        .on('click', function (d) {
          console.log('click')
          console.log(d)
        })
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
    this.renderDag()
  }
}
</script>
