import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import WorkflowExecutionList from '@/components/WorkflowExecutionList'
import WorkflowExecutionDetail from '@/components/WorkflowExecutionDetail'
import TaskExecutionLog from '@/components/TaskExecutionLog'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/helloworld',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/',
      name: 'WorkflowExecutionList',
      component: WorkflowExecutionList
    },
    {
      path: '/wf/execs/:id',
      name: 'WorkflowExecutionDetail',
      component: WorkflowExecutionDetail,
      props: true
    },
    {
      path: '/task/log/:taskId',
      name: 'TaskExecutionLog',
      component: TaskExecutionLog,
      props: true
    }
  ]
})
