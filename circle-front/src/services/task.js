// @ts-ignore
import request from '@/utils/request';
import {
  async
} from 'q';

export async function getTask(taskid, baseapi) {
  return request(baseapi + "/task/one/" + taskid, {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('token')
    }
  })
}

export async function UpdateTask(baseapi, data) {
  return request(baseapi + '/task/one/' + data.tid, {
    method: 'put',
    data: data,
    requestType: 'form',
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('token')
    }
  });
}
