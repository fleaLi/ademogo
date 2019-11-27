import {
  getTask,
  UpdateTask
} from '@/services/task';
import defaultSettings from '../../config/defaultSettings';
import {
  routerRedux
} from 'dva/router'
export default {
  namespace: 'edittask',
  state: {
    id: '',
    title: '',
    subtitle: '',
    icon: '',
    count: '',
    begin: '',
    end: '',
    limitdistance: 0,
    templateid: 0,
    count: 0,
    template: '',
    templates: [],
  },
  reducers: {
    taskrender(state, {
      payload
    }) {
      return {
        ...state,
        ...payload.data
      }
    }
  },
  effects: {
    * gettask({
      payload
    }, {
      put,
      call
    }) {
      let tid = payload.taskId
      const resp = yield call(getTask, tid, defaultSettings.apiUrl)
      yield put({
        type: 'taskrender',
        payload: resp
      })
    },
    * update({
      payload
    }, {
      put,
      call
    }) {
      const resp = yield call(UpdateTask, defaultSettings.apiUrl, payload)
      yield put(routerRedux.push('/admin/tasks'))
    }
  },
  subscriptions: {
    setup({
      dispatch,
      history
    }) {
      history.listen(({
        pathname,
        state
      }) => {
        if (pathname === "/admin/edittask") {
          if (localStorage.getItem('token') == null) {
            dispatch(routerRedux.push('/'))
          } else {
            dispatch({
              type: "gettask",
              payload: {
                taskId: state.taskId,
                token: localStorage.getItem('token')
              }
            })
          }


        }
      })
    }
  }
}
