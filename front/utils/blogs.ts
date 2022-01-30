import {
  get, post, put, axiosDestroy,
} from './request';
import { Blog } from '../types/blog';
import { FindAllResponse, FindOneResponse } from '../types/custom_response';

export const findAll = async (query: string): Promise<FindAllResponse<Blog[]>> => get(`/blogs`);
export const findOne = async (id: number): Promise<FindOneResponse<Blog>> => get(`/blogs/${id.toString()}`);
export const update = async (blog: Blog): Promise<FindOneResponse<Blog>> => put(`/blogs/${blog.id}`, { blog });
export const destroy = async (id: number): Promise<any> => axiosDestroy<any>(`/blogs/${id}`);
