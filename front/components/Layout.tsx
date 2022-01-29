import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { Popover, Transition } from "@headlessui/react";
import { Fragment } from "react";
import { Header } from "../components/Header"
function classNames(...classes: any) {
  return classes.filter(Boolean).join(" ");
}

export const Layout: React.FC = ({ children }) => {
  return (
    <>
      { children }
    </>
  );
};
