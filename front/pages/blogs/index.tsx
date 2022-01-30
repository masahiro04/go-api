import { useState, useEffect } from "react";
import type { NextPage } from "next";
import { Layout } from "../../components/Layout";
import { findAll } from "../../utils/blogs";
import { Blog } from "../../types/blog";
import Link from "next/link";

export const Index: NextPage = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);
  useEffect(() => {
    (async () => {
      const res = await findAll("");

      setBlogs(res.response);
      console.log(res.response);
    })();
  }, []);

  return (
    <Layout>
      <div className="px-6 pb-8 mx-auto sm:px-10 sm:pb-14 sm:max-w-screen-md lg:max-w-screen-lg lg:grid-cols-3 lg:gap-y-12 lg:gap-x-8 lg:pt-6">
        <ul role="list" className="divide-y divide-gray-200">
          {blogs.map((blog) => (
            <li
              key={blog.id}
              className="relative bg-white py-5 px-4 hover:bg-gray-50 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600"
            >
              <Link href={`/blogs/${blog.id}`} passHref>
                <a>
                  <div className="flex justify-between space-x-3">
                    <div className="min-w-0 flex-1">
                      <a href="#" className="block focus:outline-none">
                        <span className="absolute inset-0" aria-hidden="true" />
                        <p className="text-sm font-medium text-gray-900 truncate">
                          {blog.title}
                        </p>
                      </a>
                    </div>
                    <time
                      dateTime={blog.createdAt.toString()}
                      className="flex-shrink-0 whitespace-nowrap text-sm text-gray-500"
                    >
                      {blog.createdAt}
                    </time>
                  </div>
                  <div className="mt-1">
                    <p className="line-clamp-2 text-sm text-gray-600">
                      {blog.body}
                    </p>
                  </div>
                </a>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </Layout>
  );
};

export default Index;
