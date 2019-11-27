import React from 'react'
import moment from 'moment'
import 'moment/locale/zh-cn'
import PictureWalls from '@/components/Kit/PictureWalls'
import {
  Form,
  Icon,
  Input,
  InputNumber,
  Button,
  Upload,
  Modal,
  DatePicker,
  Radio,
  Dragger
} from 'antd';
const {
  RangePicker
} = DatePicker
moment.locale('zh-cn')
class TaskInfo extends React.Component {

  state = {
    value: 1,

  }
  hasErrors = (fieldsError) => {
    return Object.keys(fieldsError).some(field => fieldsError[field])
  }

  gettemplates = () => {
    return ( <
      Radio.Group disabled > {
        this.props.templates.map((value, index) => {
          return <Radio key = {
            index
          }
          value = {
            value.templateId
          } > {
            value.name
          } < /Radio>

        })
      } <
      /Radio.Group>
    )
  }
  validateicon = (rule, value, callback) => {
    if (!value) {
      callback(rule.message)
    } else {
      callback()
    }


  }
  handlerSubmit = (e) => {
    e.preventDefault();
    const taskicon = this.props.form.getFieldValue('taskicon')
    if (taskicon && taskicon.fileList) {
      if (!taskicon.fileList[0].response) {
        this.props.form.setFieldsValue({
          "taskicon": taskicon.fileList[0].url
        })
      } else {
        this.props.form.setFieldsValue({
          "taskicon": taskicon.fileList[0].response.data
        })
      }


    } else {
      this.props.form.setFieldsValue({
        "taskicon": this.props.taskicon
      })
    }

    this.props.form.validateFields({
      force: true
    }, (err, values) => {
      if (!err) {
        let btime = this.props.form.getFieldValue('btime')
        if (btime) {
          let ab = btime.map((value, idx) => {
            return value.format("YYYY-MM-DD HH:MM:ss")
          })
          values.btime = ab
        }
        values.tid = this.props.id
        this.props.dispatch({
          type: 'edittask/update',
          payload: values
        })
      } else {
        return err
      }

    })
  }

  render() {
      const {
        getFieldDecorator,
        getFieldsError,
        getFieldError,
        isFieldTouched
      } = this.props.form
      const ts = this.gettemplates()
      return ( < Form layout = "horizontal"
        onSubmit = {
          this.handlerSubmit
        } >
        <
        Form.Item label = "任务标题" > {
          getFieldDecorator('title', {
            initialValue: this.props.title,
            rules: [{
              required: true,
              message: '请输入任务标题!'
            }],
          })( <
            Input / > ,
          )
        }

        <
        /Form.Item> <
        Form.Item label = "任务副标题" > {
          getFieldDecorator('subtitle', {
            initialValue: this.props.subtitle,
            rules: [{
              required: true,
              message: '请输入副标题'
            }],
          })( < Input placeholder = "任务副标题" / > )
        } <
        /Form.Item> <
        Form.Item label = "任务图标" > {
          getFieldDecorator('taskicon', {
            valuePropName: "fileList",
            initialValue: this.props.taskicon ? [{
              uid: Math.random(),
              url: this.props.taskicon,
              status: 'done'
            }] : [],
            rules: [{
              validator: this.validateicon,
              required: true,
              message: "图标不能为空"
            }]
          })( < PictureWalls / > )
        } < /Form.Item> <Form.Item label = "限制距离(/m)
      " > {
      getFieldDecorator('limitdistance', {
        initialValue: this.props.limitdistance
      })( < InputNumber min = {
          1
        }
        placeholder = "限制距离(/m)" / > )
    } <
    /Form.Item> 

    <
    Form.Item label = "数量" > {
      getFieldDecorator('limitsize', {
        initialValue: this.props.limitsize,
        rules: [{
          required: true,
          message: "请输入数量"
        }]
      })( < InputNumber min = {
          0
        }
        placeholder = "数量" / > )

    } <
    /Form.Item> <
  Form.Item label = "任务有效期" > {
      getFieldDecorator('btime', {
        initialValue: [moment(this.props.beginTime), moment(this.props.endTime)],
        rules: [{
          required: true,
          message: "日期不能为空"
        }]

      })( < RangePicker showTime / > )
    } <
    /Form.Item> <
  Form.Item label = "价格" > {
      getFieldDecorator('price', {
        initialValue: this.props.price,
        rules: [{
          required: true,
          message: "请输入价格"
        }]
      })( < InputNumber min = {
          0
        }
        step = {
          0.1
        }
        placeholder = "价格" / > )
    } <
    /Form.Item> <
  Form.Item label = "模板" > {
      getFieldDecorator('templateid', {
        initialValue: this.props.template.templateId != 'undefined' ?
          this.props.template.templateId : null,
        rules: [{
          required: true,
          message: "模板不能为空"
        }]
      })(ts)

    }

    <
    /Form.Item>  <
  Form.Item >
    <
    Button type = "primary"
  htmlType = "submit"
  disabled = {
    this.hasErrors(getFieldsError)
  } > 更新 < /Button> < /
  Form.Item > <
    /Form>)
}
}

const EditTaskInfo = Form.create()(TaskInfo)

export default EditTaskInfo
