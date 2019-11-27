import request from '@/utils/request';
import TaskOperation from '@/components/Kit/TaskOperation'
import {
  Tag,
  Alert
} from 'antd';
import defaultSettings from '../../config/defaultSettings'
import {
  routerRedux
} from 'dva/router'
import editTask from './editTask';
export default {
  namespace: 'tasktable',
  state: {
    dataSource: [],
    token: localStorage.getItem('token') || '',
    page: {
      pageSize: 15,
      current: 1,
      total: 0,
      data: []
    },
    columns: [{
        title: "任务名",
        dataIndex: 'title',
        key: 'title'
      },
      {
        title: '任务状态',
        dataIndex: "taskStatus",
        key: "taskStatus",
        render: (text, record, index) => {
          if (text == 1) {
            return <Tag color = "lime" > 下线 < /Tag>
          } else if (text == 2) {
            return <Tag color = "green" > 上线 < /Tag>

          } else {
            return "其他"
          }
        },
      },
      {
        title: '开始时间',
        dataIndex: "beginTime",
        key: "beginTime"
      },
      {
        title: '结束时间',
        dataIndex: "endTime",
        key: "endTime"
      },
      {
        title: "创建时间",
        dataIndex: "creatorTime",
        key: "creatorTime"
      },
      {
        title: "操作",
        dataIndex: "operation",
        key: "operation",
        render: (text, record) => {
          return <TaskOperation src = {
            record.taskStatus
          }
          taskId = {
            record.id
          }
          />
        }
      },
    ]
  },
  reducers: {
    xr(state, {
      payload
    }) {
      const dataSource = [{
        key: '4',
        title: '胡彦斌',
        taskStatus: 5,
        creatorTime: '西湖区湖底公园1号',
      }];
      return {
        ...state,
        dataSource: dataSource
      }
    },
    netrender(state, {
      payload
    }) {
      return {
        ...state,
        page: payload.data
      }
    },
    changeone(state, {
      payload
    }) {
      const data = state.page.data.map(a => {
        if (a.id == payload.taskId) {
          a.taskStatus = payload.atarget
        }
        return a
      }).filter(a => {
        return a.taskStatus != 3
      })
      state.page.data = data
      return {
        ...state,
        ...state.page,
        data: data
      }
    }
  },
  effects: {
    * editTask({
      payload
    }, {
      put,
      call
    }) {
      let {
        taskId
      } = payload
      yield put(routerRedux.push('/admin/edittask', {
        taskId: taskId
      }))
    },
    * firstItems({
      payload
    }, {
      put,
      call
    }) {
      let {
        pageno = 1, pagesize = 15, key = '', taskstatus = '', token = ''
      } = payload
      if (!key) {
        key = ""
      }
      const baseurl = defaultSettings.apiUrl + "/task/items?pageno=" + pageno + "&pagesize=" + pagesize + "&key=" + key + "&taskstatus=" + taskstatus
      const rsp = yield call(request, baseurl, {
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
      });
      yield put({
        type: 'netrender',
        payload: rsp
      })
    },
    * taskChange({
      payload
    }, {
      put,
      call
    }) {
      let {
        taskId,
        src,
        atarget
      } = payload
      const baseurl = defaultSettings.apiUrl + "/task/change/" + taskId + "?src=" + src + "&target=" + atarget
      const rsp = yield call(request, baseurl, {
        method: "put",
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
      })
      var atid = taskId
      if (rsp.code != 200) {
        atid = ""
      }
      yield put({
        type: "changeone",
        payload: {
          taskId: atid,
          atarget: atarget
        }
      })
    },
  },
  subscriptions: {
    setup({
      dispatch,
      history
    }) {
      history.listen(({
        pathname
      }) => {
        if (pathname === "/admin/tasks") {
          if (localStorage.getItem('token') == null) {
            dispatch(routerRedux.push('/'))
          } else {
            dispatch({
              type: "firstItems",
              payload: {
                pageno: 1,
                pagesize: 15,
              }
            })
          }
        }
      })
    }
  }

}
