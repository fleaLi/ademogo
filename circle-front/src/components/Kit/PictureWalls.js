import React from 'react'
import {
  Upload,
  Icon,
  Modal
} from 'antd';

function getBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}

class PicturesWall extends React.Component {
  state = {
    previewVisible: false,
    previewImage: '',
    fileList: this.props.fileList,
  };

  handleCancel = () => this.setState({
    previewVisible: false
  });

  handlePreview = async file => {
    if (!file.url && !file.preview) {
      file.preview = await getBase64(file.originFileObj);
    }

    this.setState({
      previewImage: file.url || file.preview,
      previewVisible: true,
    });
  };

  handleChange = (e) => {
    if (e.file.status = 'done') {
      let {
        onChange
      } = this.props
      if (onChange) {
        onChange(e)

      }
    }
    this.setState({
      fileList: e.fileList
    });


  }

  render() {
    const {
      previewVisible,
      previewImage
    } = this.state;
    const uploadButton = ( < div >
      <
      Icon type = "plus" / >
      <
      div className = "ant-upload-text" > Upload < /div> </div >
    );
    const token = localStorage.getItem('token')
    const url = "/api/media/img/upload?token=" + token
    return ( < div className = "clearfix" >
      <
      Upload action = {
        url
      }
      listType = "picture-card"
      fileList = {
        this.state.fileList
      }

      onPreview = {
        this.handlePreview
      }
      onChange = {
        this.handleChange
      }
      accept = ".jpg,.png,.jpeg" > {
        this.state.fileList.length > 0 ? '' : uploadButton
      } < /Upload> <Modal visible = {
      previewVisible
    }
    footer = {
      null
    }
    onCancel = {
        this.handleCancel
      } >
      <
      img alt = "example"
    style = {
      {
        width: '100%'
      }
    }
    src = {
      previewImage
    }
    /> </Modal > < /div>
  );
}
}

export default PicturesWall
