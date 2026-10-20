function Layout({ sidebar, children }) {
  return (
    <div className="flex w-full h-screen overflow-hidden">
      <aside className="flex-shrink-0">
        {sidebar}
      </aside>
      <main className="flex-1 overflow-y-auto bg-gradient-to-br from-slate-50 to-slate-100 p-8 scrollbar-thin scrollbar-thumb-slate-300 scrollbar-track-slate-100">
        {children}
      </main>
    </div>
  );
}

export default Layout;
