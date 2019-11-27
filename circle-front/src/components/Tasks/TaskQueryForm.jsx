import { Component } from 'react';
import { Form, Select, Input, Button } from 'antd';

class TaskQueryForm extends Component {
  handlerSubmit = e => {
    e.preventDefault();

    const taskStatus = this.props.form.getFieldValue('taskStatus');
    const taskkey = this.props.form.getFieldValue('taskkey');
    this.props.dispatch({
      type: 'tasktable/firstItems',
      payload: { taskstatus: taskStatus, key: taskkey },
    });
  };

  render() {
    const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = this.props.form;

    return (
      <Form layout="inline" onSubmit={this.handlerSubmit}>
        <Form.Item>
          {getFieldDecorator('taskStatus', { initialValue: this.props.taskstatus || '' })(
            <Select>
              <Select.Option value="" key="0">
                全部
              </Select.Option>
              <Select.Option value="1" key="1">
                下线
              </Select.Option>
              <Select.Option value="2" key="2">
                上线
              </Select.Option>
            </Select>,
          )}
        </Form.Item>
        <Form.Item>
          {getFieldDecorator('taskkey')(<input placeholder="请输入任务名字" />)}
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            查询
          </Button>
        </Form.Item>
      </Form>
    );
  }
}
export default Form.create()(TaskQueryForm);
