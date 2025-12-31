import { useState } from "react";
import Navbar from "./components/Navbar";
import AuthModal from "./components/AuthModal";
import TodoList from "./components/TodoList";

export default function App() {
  const [showAuth, setShowAuth] = useState(false);

  return (
    <>
      <Navbar onLoginClick={() => setShowAuth(true)} />
      <TodoList />
      {showAuth && <AuthModal onClose={() => setShowAuth(false)} />}
    </>
  );
}
