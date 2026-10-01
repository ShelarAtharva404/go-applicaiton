import { useEffect, useState } from "react";
import api from "../services/api";
import Layout from "../components/Layout";
import "../management.css";

function Maintenance() {
  const [items, setItems] = useState([]);
  const [error, setError] = useState("");
  const user = JSON.parse(localStorage.getItem("user") || "{}");
  const canCreate = user.role === "ADMIN" || user.role === "MANAGER";

  const [form, setForm] = useState({
    vehicle_id: "",
    type: "SERVICE",
    description: "",
    cost: "",
    status: "PENDING",
  });

  const loadMaintenance = async () => {
    try {
      const token = localStorage.getItem("token");

      const res = await api.get("/maintenance", {
        headers: { Authorization: `Bearer ${token}` },
      });

      setItems(res.data || []);
    } catch (err) {
      setError(
        err.response?.data?.error ||
          "Failed to load maintenance"
      );
    }
  };

  useEffect(() => {
    loadMaintenance();
  }, []);

  const createMaintenance = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");

      await api.post(
        "/maintenance",
        {
          ...form,
          vehicle_id: Number(form.vehicle_id),
          cost: Number(form.cost || 0),
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setForm({
        vehicle_id: "",
        type: "SERVICE",
        description: "",
        cost: "",
        status: "PENDING",
      });

      await loadMaintenance();
    } catch (err) {
      setError(
        err.response?.data?.error ||
          "Failed to create maintenance"
      );
    }
  };

  return (
    <Layout>
      <div className="management-page">
        <div className="page-header">
          <div>
            <h1>Maintenance</h1>
            <p>Vehicle maintenance records</p>
          </div>
        </div>

        {error && <div className="management-error">{error}</div>}

        {canCreate && (
          <div className="form-card">
            <h2>Schedule Maintenance</h2>

            <form onSubmit={createMaintenance}>
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
                    <option value="SERVICE">SERVICE</option>
                    <option value="REPAIR">REPAIR</option>
                    <option value="INSPECTION">INSPECTION</option>
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

                <div>
                  <label>Cost</label>
                  <input
                    type="number"
                    value={form.cost}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        cost: e.target.value,
                      })
                    }
                  />
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
                    <option value="PENDING">PENDING</option>
                    <option value="IN_PROGRESS">IN_PROGRESS</option>
                    <option value="COMPLETED">COMPLETED</option>
                  </select>
                </div>
              </div>

              <button className="primary-button">
                Create Maintenance
              </button>
            </form>
          </div>
        )}

        <div className="table-card">
          <div className="table-header">
            <h2>Maintenance Records</h2>
            <span>{items.length} records</span>
          </div>

          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Vehicle</th>
                  <th>Type</th>
                  <th>Description</th>
                  <th>Cost</th>
                  <th>Status</th>
                </tr>
              </thead>

              <tbody>
                {items.map((item) => (
                  <tr key={item.id}>
                    <td>#{item.id}</td>
                    <td>{item.vehicle_id}</td>
                    <td>{item.type}</td>
                    <td>{item.description}</td>
                    <td>₹{item.cost}</td>
                    <td>{item.status}</td>
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

export default Maintenance;
