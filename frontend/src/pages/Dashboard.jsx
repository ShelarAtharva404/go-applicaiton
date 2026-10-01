import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";
import Layout from "../components/Layout";
import "../dashboard.css";

function Dashboard() {
  const navigate = useNavigate();

  const [summary, setSummary] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const user = JSON.parse(localStorage.getItem("user") || "{}");

  useEffect(() => {
    const fetchDashboard = async () => {
      try {
        const token = localStorage.getItem("token");

        const response = await api.get("/dashboard/summary", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        setSummary(response.data);
      } catch (err) {
        if (err.response?.status === 401) {
          localStorage.removeItem("token");
          localStorage.removeItem("user");
          navigate("/");
          return;
        }

        setError("Failed to load dashboard");
      } finally {
        setLoading(false);
      }
    };

    fetchDashboard();
  }, [navigate]);

  if (loading) {
    return (
      <div className="dashboard-loading">
        Loading FleetPulse...
      </div>
    );
  }

  if (error) {
    return (
      <div className="dashboard-loading">
        {error}
      </div>
    );
  }

  return (
    <Layout>
      <div className="dashboard">
        <header className="dashboard-header">
          <div>
            <h1>
              Fleet<span>Pulse</span>
            </h1>

            <p>Fleet Management Dashboard</p>
          </div>

          <div className="user-section">
            <div>
              <strong>{user.name}</strong>
              <small>{user.role}</small>
            </div>
          </div>
        </header>

        <main className="dashboard-content">
          <h2>Overview</h2>

          <div className="stats-grid">
            <div className="stat-card">
              <h3>Vehicles</h3>
              <div className="stat-number">
                {summary.vehicles.total}
              </div>
              <p>
                {summary.vehicles.active} active
              </p>
            </div>

            <div className="stat-card">
              <h3>Drivers</h3>
              <div className="stat-number">
                {summary.drivers.total}
              </div>
              <p>
                {summary.drivers.active} active
              </p>
            </div>

            <div className="stat-card">
              <h3>Trips</h3>
              <div className="stat-number">
                {summary.trips.total}
              </div>
              <p>
                {summary.trips.planned} planned
              </p>
            </div>

            <div className="stat-card">
              <h3>Incidents</h3>
              <div className="stat-number">
                {summary.incidents.open}
              </div>
              <p>open incidents</p>
            </div>

            <div className="stat-card">
              <h3>Maintenance</h3>
              <div className="stat-number">
                {summary.maintenance.pending}
              </div>
              <p>pending maintenance</p>
            </div>

            <div className="stat-card">
              <h3>Completed Trips</h3>
              <div className="stat-number">
                {summary.trips.completed}
              </div>
              <p>completed trips</p>
            </div>
          </div>
        </main>
      </div>
    </Layout>
  );
}

export default Dashboard;
