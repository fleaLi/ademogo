import request from '@/utils/request';
import {
  async
} from 'q';
export async function query() {
  return request('/api/users');
}
export async function queryCurrent() {
  return request('/api/currentUser');
}

export async function login(data, baseapi) {
  return request(baseapi + '/login', {
    method: 'post',
    data: data,
    requestType: 'form',
  });

}
