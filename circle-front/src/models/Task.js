import request from '@/utils/request'
import {
  routerRedux
} from 'dva/router'
import defaultSettings from '../../config/defaultSettings'

export default {
  namespace: "task",
  state: {
    templates: [{
        id: 1,
        name: "模板1"
      },
      {
        id: 2,
        name: "模板2"
      },
    ],
  },
  reducers: {
    aftersub(state, {
      payload
    }) {
      if (payload.code == 200) {
        //跳转到tasks
      } else {
        //修改state
      }
    },
    constinfo(state, {
      payload
    }) {
      return {
        ...state,
        ...payload.data
      }
    },
    showinfo(state, {
      payload
    }) {
      return {
        ...state,
        ...payload
      }
    },
  },
  effects: {
    * create({
      payload
    }, {
      put,
      call
    }) {
      const baseurl = defaultSettings.apiUrl + "/task/task"
      try {
        const resp = yield call(request, baseurl, {
          method: 'post',
          requestType: "form",
          data: payload,
          headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
          }
        })
        if (resp.status == 200) {
          payload.error = resp.msg
          yield put({
            type: "showinfo",
            payload: payload
          })
        } else {
          yield put(routerRedux.push('/admin/tasks'))

        }

      } catch (e) {
        payload.error = e
        yield put({
          type: "showinfo",
          payload: payload
        })
      }


    },
    * getInfo({
      payload
    }, {
      put,
      call
    }) {
      const baseurl = defaultSettings.apiUrl + "/task/constinfo"
      const resp = yield call(request, baseurl, {
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
      })
      yield put({
        type: "constinfo",
        payload: resp
      })
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
        if (pathname === "/admin/pubtask") {
          if (localStorage.getItem('token') == null) {
            dispatch(routerRedux.push('/'))
          } else {
            dispatch({
              type: "getInfo",
              payload: {
                pageno: 1,
                pagesize: 15,
                token: localStorage.getItem('token')
              }
            })
          }


        }
      })
    }
  }

}
