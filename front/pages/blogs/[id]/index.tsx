import { useState, useEffect } from "react";

import type { NextPage } from "next";
import { Layout } from "../../../components/Layout";
import { findAll, findOne } from "../../../utils/blogs";
import { Blog } from "../../../types/blog";
import { useRouter } from "next/router";
import Link from "next/link";

const messages = [
  {
    id: 1,
    subject: "Velit placeat sit ducimus non sed",
    sender: "Gloria Roberston",
    time: "1d ago",
    datetime: "2021-01-27T16:35",
    preview:
      "Doloremque dolorem maiores assumenda dolorem facilis. Velit vel in a rerum natus facere. Enim rerum eaque qui facilis. Numquam laudantium sed id dolores omnis in. Eos reiciendis deserunt maiores et accusamus quod dolor.",
  },
];

export const Index: NextPage = () => {
  const [blog, setBlog] = useState<Blog>();
  const router = useRouter();
  const { id } = router.query;

  useEffect(() => {
    (async () => {
      const res = await findOne(Number(id));
      setBlog(res.response);
    })();
  }, []);

  return (
    <Layout>
      <div className="px-6 pb-8 mx-auto sm:px-10 sm:pb-14 sm:max-w-screen-md lg:max-w-screen-lg lg:grid-cols-3 lg:gap-y-12 lg:gap-x-8 lg:pt-6">
        <Link href={`/blogs/${id}/edit`}>
          <a>
            <button
              type="button"
              className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              編集
            </button>
          </a>
        </Link>{" "}
        <button
          type="button"
          className="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
        >
          削除
        </button>
        <ul role="list" className="divide-y divide-gray-200">
          {messages.map((message) => (
            <li
              key={message.id}
              className="relative bg-white py-5 px-4 hover:bg-gray-50 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600"
            >
              <div className="flex justify-between space-x-3">
                <div className="min-w-0 flex-1">
                  <a href="#" className="block focus:outline-none">
                    <span className="absolute inset-0" aria-hidden="true" />
                    <p className="text-sm font-medium text-gray-900 truncate">
                      {message.sender}
                    </p>
                    <p className="text-sm text-gray-500 truncate">
                      {message.subject}
                    </p>
                  </a>
                </div>
                <time
                  dateTime={message.datetime}
                  className="flex-shrink-0 whitespace-nowrap text-sm text-gray-500"
                >
                  {message.time}
                </time>
              </div>
              <div className="mt-1">
                <p className="line-clamp-2 text-sm text-gray-600">
                  {message.preview}
                </p>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </Layout>
  );
};

export default Index;
