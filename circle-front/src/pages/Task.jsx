import {connect} from 'dva'
import TaskInfo from '@/components/Task/TaskInfo'


export default connect(({settings,task})=>({...settings,...task}))(TaskInfo)