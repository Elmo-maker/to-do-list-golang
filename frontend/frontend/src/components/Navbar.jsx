export default function Navbar({ onLoginClick }) {
  return (
    <nav className="flex items-center justify-between px-6 py-4 bg-slate-900 border-b border-slate-800">
      <h1 className="text-xl font-bold text-white">Todo App</h1>

      <button
        onClick={onLoginClick}
        className="rounded-lg bg-blue-600 px-5 py-2 font-medium text-white hover:bg-blue-700 transition"
      >
        Login
      </button>
    </nav>
  );
}
