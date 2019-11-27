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
  Dragger,
  message
} from 'antd';
const {
  RangePicker
} = DatePicker
moment.locale('zh-cn')
class TaskInfo extends React.Component {

    state = {
      value: 1,

    }
    handlerupload = (e) => {
      let fileList = e.fileList.map(file => {
        if (file.response) {
          if (file.response.code != 200) {
            file.url = file.response.data
            return file
          }
        }
        return file;
      })
      return fileList
    }
    normFile = e => {
      if (Array.isArray(e)) {
        return e;
      }
      return e && e.fileList;
    };
    hasErrors = (fieldsError) => {
      return Object.keys(fieldsError).some(field => fieldsError[field])
    }

    gettemplates = () => {
      return ( < Radio.Group > {
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
        } < /Radio.Group>)
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
          this.props.form.setFieldsValue({
            "taskicon": taskicon.fileList[0].response.data
          })

        }



        this.props.form.validateFields({
          force: true
        }, (err, values) => {
          if (!err) {
            let pointsref = this.props.form.getFieldValue('pointsref')
            let ps = pointsref.filter((value) => {
              return value.status == 'done' && value.uid != '-1'
            }).map((value, idx) => {
              return value.response.data
            });
            if (ps.length < 1) {
              return message.warn("点位文件不能为空！")
            }
            values.pointsref = ps
            let btime = this.props.form.getFieldValue('btime')
            if (btime) {
              let ab = btime.map((value, idx) => {
                return value.format("YYYY-MM-DD HH:MM:ss")
              })
              values.btime = ab
            }
            this.props.dispatch({
              type: 'task/create',
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
        if (this.props.error) {
          message.error(this.props.error)
        }
        const ts = this.gettemplates()
        return ( < Form layout = "horizontal"
          onSubmit = {
            this.handlerSubmit
          } >
          <
          Form.Item label = "任务标题" > {
            getFieldDecorator('title', {
              rules: [{
                required: true,
                message: '请输入任务标题!'
              }],
            })( <
              Input placeholder = "如：xxx广告任务" / > ,
            )
          }

          <
          /Form.Item> <Form.Item label = "任务副标题" > {
          getFieldDecorator('subtitle', {
            initialValue: this.props.subtitle,
            rules: [{
              required: true,
              message: '请输入副标题'
            }],
          })( < Input.TextArea placeholder = "可以输入执行城市" / > )
        } < /Form.Item> <Form.Item label = "任务图标" > {
        getFieldDecorator('taskicon', {
          valuePropName: "fileList",
          initialValue: this.props.taskicon ? [{
            uid: '-1',
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
  } < /Form.Item> 

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

  } < /Form.Item> <Form.Item label = "任务有效期" > {
getFieldDecorator('btime', {
  rules: [{
    required: true,
    message: "日期不能为空"
  }]

})( < RangePicker showTime / > )
} < /Form.Item> <Form.Item label = "价格" > {
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
} < /Form.Item> <Form.Item label = "模板" > {
getFieldDecorator('templateid', {
  rules: [{
    required: true,
    message: "模板不能为空"
  }]
})(ts)

}

<
/Form.Item> <Form.Item label = "点位信息" > <
div className = "dropbox" > {
    getFieldDecorator('pointsref', {
      valuePropName: 'fileList',
      initialValue: [{
        uid: '-1',
        name: '样例',
        status: 'done',
        url: 'http://antzb-production.oss-cn-qingdao.aliyuncs.com/qa/example0811.xlsx'
      }],
      getValueFromEvent: this.normFile,
      rules: [{
        validator: this.validateicon,
        required: true,
        message: "点位不能为空"
      }]
    })( < Upload.Dragger accept = ".xls,.xlsx"
      multiple = {
        false
      }
      onChange = {
        this.handlerupload
      }

      name = "files"
      action = {
        this.props.apiUrl + "/media/csv/upload?token=" + localStorage.getItem('token')
      } >
      <
      p className = "ant-upload-drag-icon" >
      <
      Icon type = "inbox" / >
      <
      /p> <p className = "ant-upload-text" > 点击或者拖拽即可上传 </p > < p className = "ant-upload-hint" > 支持单个或多个 < /p>  </Upload.Dragger > )

  } < /div> </Form.Item > < Form.Item >
  <
  Button type = "primary"
htmlType = "submit"
disabled = {
  this.hasErrors(getFieldsError)
} > 发布 < /Button> </Form.Item > < /Form>)
}
}

const TaskInfo1 = Form.create()(TaskInfo)

export default TaskInfo1
