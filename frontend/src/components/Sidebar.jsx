function Sidebar({ onNavigate }) {
  return (
    <div className="w-64 h-screen bg-gradient-to-b from-slate-800 to-slate-900 text-white flex flex-col shadow-2xl">
      <div className="p-5 border-b border-white/10">
        <h2 className="text-2xl font-semibold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
          PowerMoat
        </h2>
      </div>
      
      <nav className="flex-1 py-3 overflow-y-auto">
      </nav>
    </div>
  );
}

export default Sidebar;
