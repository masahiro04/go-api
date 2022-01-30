export interface Blog {
  id: number;
  title: string;
  body: string;
  createdAt: Date;
  updatedAt: Date;
}

export type BlogForm = Pick<Blog, 'title' | 'body'>;
