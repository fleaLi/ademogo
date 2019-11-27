import { connect } from 'dva';
import TaskInfo from '@/components/Task/EditTaskInfo';

export default connect(({ settings, edittask }) => ({ ...settings, ...edittask }))(TaskInfo);
