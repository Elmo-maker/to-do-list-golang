import { useState } from "react";
import { createPortal } from "react-dom";

export default function AuthModal({ onClose }) {
  const [isLogin, setIsLogin] = useState(true);

  return createPortal(
    <div className="fixed inset-0 z-[9999] flex items-center justify-center bg-black/70 backdrop-blur-sm">
      <div className="relative w-96 rounded-2xl bg-slate-800 border border-slate-700 shadow-2xl p-6 animate-scaleIn">
        
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-slate-400 hover:text-white"
        >
          ✕
        </button>

        <h2 className="text-2xl font-bold text-center mb-6 text-white">
          {isLogin ? "Welcome Back 👋" : "Create Account 🚀"}
        </h2>

        <input
          type="email"
          placeholder="Email"
          className="w-full mb-3 rounded-lg bg-slate-900 border border-slate-700 p-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full mb-5 rounded-lg bg-slate-900 border border-slate-700 p-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
        />

        <button className="w-full rounded-lg bg-blue-600 py-3 font-semibold text-white hover:bg-blue-700 transition">
          {isLogin ? "Login" : "Register"}
        </button>

        <p className="mt-5 text-center text-sm text-slate-400">
          {isLogin ? "Belum punya akun?" : "Sudah punya akun?"}{" "}
          <span
            onClick={() => setIsLogin(!isLogin)}
            className="cursor-pointer text-blue-400 hover:underline"
          >
            {isLogin ? "Register" : "Login"}
          </span>
        </p>
      </div>
    </div>,
    document.body
  );
}
