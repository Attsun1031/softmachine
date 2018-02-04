<template>
  <div>
    <h3>WorkflowExecution List</h3>
    <div>
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
          <td>{{ msg }}</td>
          <td>Eclair</td>
          <td>$0.87</td>
          <td>$0.87</td>
          <td>$0.87</td>
        </tr>
        <tr>
          <td>{{ msg }}</td>
          <td>Eclair</td>
          <td>$0.87</td>
          <td>$0.87</td>
          <td>$0.87</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'WorkflowExecutionList',
  data () {
    return {
      msg: 'test'
    }
  }
}
</script>
