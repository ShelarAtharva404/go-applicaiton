import { NavLink, useNavigate } from "react-router-dom";
import "../layout.css";

function Layout({ children }) {
  const navigate = useNavigate();

  const user = JSON.parse(localStorage.getItem("user") || "{}");

  const logout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    navigate("/");
  };

  return (
    <div className="app-layout">
      <aside className="sidebar">
        <div className="sidebar-logo">
          Fleet<span>Pulse</span>
        </div>

        <div className="sidebar-subtitle">
          Fleet Management
        </div>

        <nav className="sidebar-nav">
          <NavLink to="/dashboard">
            Dashboard
          </NavLink>

          <NavLink to="/vehicles">
            Vehicles
          </NavLink>

          <NavLink to="/drivers">
            Drivers
          </NavLink>

          <NavLink to="/trips">
            Trips
          </NavLink>

          <NavLink to="/maintenance">
            Maintenance
          </NavLink>

          <NavLink to="/incidents">
            Incidents
          </NavLink>
        </nav>

        <div className="sidebar-user">
          <div className="sidebar-user-name">
            {user.name}
          </div>

          <div className="sidebar-user-role">
            {user.role}
          </div>

          <button onClick={logout}>
            Logout
          </button>
        </div>
      </aside>

      <main className="app-main">
        {children}
      </main>
    </div>
  );
}
export default Layout;
