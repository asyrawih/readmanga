import Link from "next/link"

export const Sidebar = () => {
  return (
    <aside className="hidden sm:block h-screen w-1/6 text-white border">
      <div className="flex flex-col ml-4 mt-3">
        <span className="text-3xl">Dashboard</span>
        <ul className="mt-12">
          <li className="mt-3 p-3 cursor-pointer">Chapters</li>
          <li className="mt-3 p-3 cursor-pointer">Posts</li>
          <Link href={`/dashboard/manga`}>
            <li className="mt-3 p-3 cursor-pointer">Manga</li>
          </Link>
          <li className="mt-3 p-3 cursor-pointer">Members</li>
          <li className="mt-3 p-3 cursor-pointer">Users</li>
        </ul>
      </div>
    </aside>
  )
}
