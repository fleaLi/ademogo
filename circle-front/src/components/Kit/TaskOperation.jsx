import { Component } from 'react';
import { Menu, Dropdown, Tag } from 'antd';
import { connect } from 'dva';
class TaskOperation extends Component {
  menuclick = ({ item, key, keypath, domevent }) => {
    if ('上线' == key) {
      this.props.dispatch({
        type: 'tasktable/taskChange',
        payload: { taskId: this.props.taskId, src: this.props.src, atarget: 2 },
      });
    } else if ('下线' == key) {
      this.props.dispatch({
        type: 'tasktable/taskChange',
        payload: { taskId: this.props.taskId, src: this.props.src, atarget: 1 },
      });
    } else if ('删除' == key) {
      this.props.dispatch({
        type: 'tasktable/taskChange',
        payload: { taskId: this.props.taskId, src: this.props.src, atarget: 3 },
      });
    } else if ('编辑' == key) {
      this.props.dispatch({
        type: 'tasktable/editTask',
        payload: { taskId: this.props.taskId },
      });
    }
  };
  render() {
    const { taskId, src } = this.props;
    var menu;
    if (src == 1) {
      menu = (
        <Menu onClick={this.menuclick}>
          <Menu.Item key="编辑">
            {' '}
            <Tag>编辑</Tag>{' '}
          </Menu.Item>
          <Menu.Item key="上线">
            {' '}
            <Tag color="green">上线</Tag>{' '}
          </Menu.Item>{' '}
          <Menu.Item key="删除">
            {' '}
            <Tag color="red">删除</Tag>{' '}
          </Menu.Item>
        </Menu>
      );
    } else if (src == 2) {
      menu = (
        <Menu onClick={this.menuclick}>
          <Menu.Item key="下线">
            {' '}
            <Tag>下线</Tag>{' '}
          </Menu.Item>{' '}
        </Menu>
      );
    } else {
      menu = (
        <Menu>
          <Menu.Item>
            {' '}
            <Tag>暂无操作</Tag>{' '}
          </Menu.Item>{' '}
        </Menu>
      );
    }

    return (
      <Dropdown overlay={menu}>
        <Tag color="blue">操作</Tag>
      </Dropdown>
    );
  }
}

export default connect(({ tasktable }) => ({
  ...tasktable,
}))(TaskOperation);
