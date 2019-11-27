import React from 'react';
import { Form, Icon, Input, Button, Checkbox, Layout, Result, Row, Col } from 'antd';
const { Header, Footer, Content } = Layout;
import { routerRedux } from 'dva/router';
import { connect } from 'dva';
class NormalLoginForm extends React.Component {
  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        this.props.dispatch({ type: 'user/dvalogin', payload: values });
      }
    });
  };
  relogin = () => {
    this.props.message = null;
    routerRedux.push('/');
  };
  render() {
    const { getFieldDecorator } = this.props.form;
    if (this.props.message) {
      return (
        <Result
          title={this.props.message}
          icon={<Icon type="frown" theme="twoTone"></Icon>}
          extra={
            <Button type="primary" onClick={this.relogin}>
              返回
            </Button>
          }
        ></Result>
      );
    }
    return (
      <div>
        <Row type="flex" justify="space-around" align="middle">
          <Col span={4}></Col>
          <Col span={4}>
            <Form onSubmit={this.handleSubmit} className="login-form">
              <Form.Item>
                {getFieldDecorator('username', {
                  rules: [{ required: true, message: '请输入用户名！' }],
                })(
                  <Input
                    prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                    placeholder="用户名"
                  />,
                )}
              </Form.Item>
              <Form.Item>
                {getFieldDecorator('password', {
                  rules: [{ required: true, message: '请输入密码!' }],
                })(
                  <Input
                    prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                    type="password"
                    placeholder="密码"
                  />,
                )}
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" className="login-form-button">
                  登录
                </Button>
              </Form.Item>
            </Form>
          </Col>
          <Col span={4}></Col>
        </Row>
      </div>
    );
  }
}
const WrappedNormalLoginForm = Form.create({ name: 'normal_login' })(NormalLoginForm);

export default connect(({ user }) => ({ ...user }))(WrappedNormalLoginForm);
