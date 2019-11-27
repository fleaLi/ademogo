import {
  Table,
  Divider,
  Form,
  Input,
  Button,
  Select
} from 'antd'
import {
  Component
} from 'react';
import {
  connect
} from 'dva'
import './index.less'
import TaskQueryForm from '@/components/Tasks/TaskQueryForm'
class TaskLists extends Component {
  constructor(props) {
    super(props)
    this.state = {
      columns: props.columns
    }
  }
  componentDidMount() {}

  render() {
    const page = {
      pageSize: this.props.page.pageSize,
      current: this.props.page.current,
      total: this.props.page.total,
      onChange: (pageno, psize) => this.props.dispatch({
        type: 'tasktable/firstItems',
        payload: {
          pagesize: psize,
          pageno: pageno,
          key: this.props.page.key,
          taskstatus: this.props.page.taskstatus
        }
      }),
    }
    return <div >
      <TaskQueryForm {
        ...this.props
      }/> <Table rowKey = {
      record => record.id
    }
    dataSource = {
      this.props.page.data
    }
    columns = {
      this.state.columns
    }
    pagination = {
        page
      }> </Table> </div>

  }
}

function mapstatetoprops(tasktable) {
  return {
    tasktable
  }
}

// export default  connect(mapstatetoprops)(TaskLists)
export default TaskLists
