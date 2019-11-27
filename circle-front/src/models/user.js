import {
  queryCurrent,
  query as queryUsers,
  login
} from '@/services/user';
import defaultSettings from '../../config/defaultSettings'
import {
  routerRedux
} from 'dva/router'

const UserModel = {
  namespace: 'user',
  state: {
    currentUser: {},
  },
  effects: {
    * dvalogin({
      payload
    }, {
      call,
      put
    }) {
      const response = yield call(login, payload, defaultSettings.apiUrl)
      if (response.code != 200) {
        yield put({
          type: "loginfailed",
          payload: response
        })
      } else {
        localStorage.setItem('token', response.token)
        localStorage.setItem('uname', payload.username)
        yield put(routerRedux.push("/admin/tasks", {
          "token": response.token
        }))
      }

    },
    * loginout({
      payload
    }, {
      put,
      call
    }) {
      localStorage.removeItem('token');
      localStorage.removeItem('uname');
      yield put(routerRedux.push('/'))
    }
  },
  reducers: {
    loginout(state, action) {
      localStorage.removeItem('token');
      localStorage.removeItem('uname');
      return {
        ...state
      }
    },
    loginfailed(state, {
      payload
    }) {
      return {
        ...state,
        ...payload.message
      }
    }

  },
};
export default UserModel;
