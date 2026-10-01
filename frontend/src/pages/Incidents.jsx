import { useEffect, useState } from "react";
import api from "../services/api";
import Layout from "../components/Layout";
import "../management.css";

function Incidents() {
  const [incidents, setIncidents] = useState([]);
  const [error, setError] = useState("");
  const user = JSON.parse(localStorage.getItem("user") || "{}");
  const canCreate = user.role === "ADMIN" || user.role === "MANAGER";

  const [form, setForm] = useState({
    vehicle_id: "",
    driver_id: "",
    type: "BREAKDOWN",
    description: "",
    severity: "LOW",
    status: "OPEN",
  });

  const loadIncidents = async () => {
    try {
      const token = localStorage.getItem("token");

      const res = await api.get("/incidents", {
        headers: { Authorization: `Bearer ${token}` },
      });

      setIncidents(res.data || []);
    } catch (err) {
      setError(
        err.response?.data?.error ||
          "Failed to load incidents"
      );
    }
  };

  useEffect(() => {
    loadIncidents();
  }, []);

  const createIncident = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");

      await api.post(
        "/incidents",
        {
          ...form,
          vehicle_id: Number(form.vehicle_id),
          driver_id: form.driver_id
            ? Number(form.driver_id)
            : null,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setForm({
        vehicle_id: "",
        driver_id: "",
        type: "BREAKDOWN",
        description: "",
        severity: "LOW",
        status: "OPEN",
      });

      await loadIncidents();
    } catch (err) {
      setError(
        err.response?.data?.error ||
          "Failed to create incident"
      );
    }
  };

  return (
    <Layout>
      <div className="management-page">
        <div className="page-header">
          <div>
            <h1>Incidents</h1>
            <p>Track vehicle incidents</p>
          </div>
        </div>

        {error && <div className="management-error">{error}</div>}

        {canCreate && (
          <div className="form-card">
            <h2>Report Incident</h2>

            <form onSubmit={createIncident}>
              <div className="form-grid">
                <div>
                  <label>Vehicle ID</label>
                  <input
                    type="number"
                    value={form.vehicle_id}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        vehicle_id: e.target.value,
                      })
                    }
                    required
                  />
                </div>

                <div>
                  <label>Driver ID</label>
                  <input
                    type="number"
                    value={form.driver_id}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        driver_id: e.target.value,
                      })
                    }
                  />
                </div>

                <div>
                  <label>Type</label>
                  <select
                    value={form.type}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        type: e.target.value,
                      })
                    }
                  >
                    <option value="BREAKDOWN">BREAKDOWN</option>
                    <option value="ACCIDENT">ACCIDENT</option>
                    <option value="DAMAGE">DAMAGE</option>
                    <option value="OTHER">OTHER</option>
                  </select>
                </div>

                <div>
                  <label>Severity</label>
                  <select
                    value={form.severity}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        severity: e.target.value,
                      })
                    }
                  >
                    <option value="LOW">LOW</option>
                    <option value="MEDIUM">MEDIUM</option>
                    <option value="HIGH">HIGH</option>
                    <option value="CRITICAL">CRITICAL</option>
                  </select>
                </div>

                <div>
                  <label>Status</label>
                  <select
                    value={form.status}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        status: e.target.value,
                      })
                    }
                  >
                    <option value="OPEN">OPEN</option>
                    <option value="INVESTIGATING">
                      INVESTIGATING
                    </option>
                    <option value="RESOLVED">RESOLVED</option>
                  </select>
                </div>

                <div>
                  <label>Description</label>
                  <input
                    value={form.description}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        description: e.target.value,
                      })
                    }
                  />
                </div>
              </div>

              <button className="primary-button">
                Report Incident
              </button>
            </form>
          </div>
        )}

        <div className="table-card">
          <div className="table-header">
            <h2>Incidents</h2>
            <span>{incidents.length} incidents</span>
          </div>

          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Vehicle</th>
                  <th>Driver</th>
                  <th>Type</th>
                  <th>Severity</th>
                  <th>Status</th>
                  <th>Description</th>
                </tr>
              </thead>

              <tbody>
                {incidents.map((incident) => (
                  <tr key={incident.id}>
                    <td>#{incident.id}</td>
                    <td>{incident.vehicle_id}</td>
                    <td>{incident.driver_id || "-"}</td>
                    <td>{incident.type}</td>
                    <td>{incident.severity}</td>
                    <td>{incident.status}</td>
                    <td>{incident.description}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default Incidents;
