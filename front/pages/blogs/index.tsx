import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

import { Header } from "../../components/Header";
import {Footer } from "../../components/Footer";

const Index: NextPage = () => {
  return (
    <>
      <Header />
      <h1 className="text-3xl font-bold underline">Blog index</h1>
      <Footer /> 
    </>
  );
};

export default Index;
