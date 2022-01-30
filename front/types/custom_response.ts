export interface FindAllResponse<T> {
  count: number;
  response: T;
}

export type FindOneResponse<T> = Pick<FindAllResponse<T>, 'response'>;

