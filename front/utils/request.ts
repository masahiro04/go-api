import axios, { AxiosInstance, AxiosResponse } from 'axios';
// import { API_PATH } from '../config/api';
// import mockAdapter from '../mocks/axios/adapter';
// import { getToken } from './localstorage';
import { KeyVal } from '../types/common';

const instance = (withCred = true): AxiosInstance => {
  const headers: KeyVal = { 'Content-Type': 'application/json' };
  if (withCred) {
    // headers.Authorization = `Token ${getToken()}`;
    headers.Authorization = `Token `;
  }

  const ax = axios.create({
    baseURL: 'http://localhost:8080/api', timeout: 15000, headers, withCredentials: withCred,
  });

  ax.interceptors.response.use((response) => response, async (error) => {
    // NOTO: sessionが切れたら実行
    // TODO: UIをいい感じにしたい。
    if (error.response.status === 401) {
      // eslint-disable-next-line no-alert
      // alert('ログインの有効期限が切れました。ログインしなおしてください。');
      // window.location.href = '/login';
      return Promise.reject(error);
    }

    return Promise.reject(error);
  });
  return ax;
};

const responseBody = (response: AxiosResponse) => response.data;

export const get = <T= unknown>(url: string, cred = true): Promise<T> =>
  instance(cred).get<T>(url).then(responseBody);

export const post = <T = unknown>(url: string, body: {}, cred = true): Promise<T> =>
  instance(cred).post<T>(url, body).then(responseBody);

export const put = <T  = unknown>(url: string, body: {}, cred = true): Promise<T> =>
  instance(cred).put<T>(url, body).then(responseBody);

export const axiosDestroy = <T = unknown>(url: string, cred = true): Promise<T> =>
  instance(cred).delete<T>(url).then(responseBody);
