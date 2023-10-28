import type { Metadata } from 'next'
import { Inter as FontSans } from "next/font/google"
import { cn } from '@/lib/utils'
import { ThemeProvider } from '@/provider/theme'
import "./globals.css"


export const metadata: Metadata = {
  title: 'Bacakomik',
  description: '',
}

const fontSans = FontSans({
  subsets: ['latin'],
  variable: '--font-sans'
})

export default function RootLayout({ children, }: { children: React.ReactNode }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={cn(
          'min-h-screen font-sans antialiased dark:bg-gray-900',
          fontSans.variable
        )}
      >
        <ThemeProvider
          attribute='class'
          defaultTheme='system'
          enableSystem
        >
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
