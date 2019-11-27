import {connect} from 'dva'
import TaskLists from '@/components/Tasks/TaskListsModal'

export default connect(({tasktable})=>({...tasktable}))(TaskLists)