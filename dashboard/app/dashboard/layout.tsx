'use client'
import { Navbar } from "@/components/navbar";
import { Sidebar } from "@/components/sidebar";
import { Toaster } from "@/components/ui/toaster";
import { ReactNode } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from 'react-query/devtools'

export default function Dashboardlayout({ children }: { children: ReactNode }) {
  const client = new QueryClient()
  return (
    <>
      <QueryClientProvider client={client}>
        <ReactQueryDevtools />
        <Navbar />
        <div className="flex">
          <Sidebar />
          <div className="flex-1 text-white mx-2">
            {children}
            <Toaster />
          </div>
        </div>
      </QueryClientProvider>
    </>

  )
}
