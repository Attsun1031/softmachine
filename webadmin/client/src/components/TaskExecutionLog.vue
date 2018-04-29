<template>
  <div>
    <h3>TaskExecution Log</h3>
    <div>
      <label for="pod-select">Pod</label>
      <select v-model="currentPodName" id="pod-select" class="browser-default">
        <option v-for="pod in pods" :key="pod.podName" :value="pod.podName">
          {{ pod.podName }}
        </option>
      </select>

      <label for="container-select">Container</label>
      <select v-if="currentPod" v-model="currentContainer" id="container-select" class="browser-default">
        <option v-for="container in currentPod.containers" :key="container" :value="container">
          {{ container }}
        </option>
      </select>

      <h5>Log</h5>
      <div class="card-panel teal lighten-5">
        {{ log }}
      </div>
    </div>
  </div>
</template>

<script>
import JobnetesApi from '@/external/jobnetesApi'

export default {
  name: 'TaskExecutionLog',
  props: ['taskId'],
  data: function () {
    return {
      pods: [],
      currentPodName: '',
      currentContainer: '',
      log: ''
    }
  },
  computed: {
    currentPod: function () {
      return this.pods.find(v => v.podName === this.currentPodName)
    }
  },
  watch: {
    currentContainer: function (val) {
      JobnetesApi.getContainerLog(this.taskId, this.currentPodName, val)
        .then(response => response.json())
        .then(data => {
          this.log = data.item
        })
    }
  },
  beforeRouteEnter: function (route, redirect, next) {
    next(vm => {
      // JobnetesApi.getPods(vm.taskId)
      JobnetesApi.getPods(vm.taskId)
        .then(response => response.json())
        .then(data => {
          vm.pods = data.items
          if (vm.pods.length > 0) {
            vm.currentPodName = vm.pods[0].podName
            if (vm.pods[0].containers.length > 0) {
              vm.currentContainer = vm.pods[0].containers[0]
            }
          }
        })
    })
  },
  beforeRouteUpdate (to, from, next) {
    this.pods = []
    this.currentContainer = ''
    this.log = ''
    JobnetesApi.getPods(this.taskId)
      .then(response => response.json())
      .then(data => {
        this.pods = data.items
        next()
      })
  }
}
</script>
