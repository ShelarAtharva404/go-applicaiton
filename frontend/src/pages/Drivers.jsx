import { useEffect, useState } from "react";
import api from "../services/api";
import Layout from "../components/Layout";
import "../management.css";

function Drivers() {
  const [drivers, setDrivers] = useState([]);
  const [error, setError] = useState("");
  const user = JSON.parse(localStorage.getItem("user") || "{}");
  const isAdmin = user.role === "ADMIN";

  const [form, setForm] = useState({
    name: "",
    email: "",
    phone: "",
    license_number: "",
    status: "ACTIVE",
    vehicle_id: "",
  });

  const loadDrivers = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await api.get("/drivers", {
        headers: { Authorization: `Bearer ${token}` },
      });
      setDrivers(res.data || []);
    } catch (err) {
      setError(err.response?.data?.error || "Failed to load drivers");
    }
  };

  useEffect(() => {
    loadDrivers();
  }, []);

  const createDriver = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");

      await api.post(
        "/drivers",
        {
          ...form,
          vehicle_id: form.vehicle_id
            ? Number(form.vehicle_id)
            : null,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setForm({
        name: "",
        email: "",
        phone: "",
        license_number: "",
        status: "ACTIVE",
        vehicle_id: "",
      });

      await loadDrivers();
    } catch (err) {
      setError(err.response?.data?.error || "Failed to create driver");
    }
  };

  const deleteDriver = async (id) => {
    if (!window.confirm("Delete this driver?")) return;

    try {
      const token = localStorage.getItem("token");

      await api.delete(`/drivers/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      await loadDrivers();
    } catch (err) {
      setError(err.response?.data?.error || "Failed to delete driver");
    }
  };

  return (
    <Layout>
      <div className="management-page">
        <div className="page-header">
          <div>
            <h1>Drivers</h1>
            <p>Manage fleet drivers</p>
          </div>
        </div>

        {error && <div className="management-error">{error}</div>}

        {isAdmin && (
          <div className="form-card">
            <h2>Add Driver</h2>

            <form onSubmit={createDriver}>
              <div className="form-grid">
                <div>
                  <label>Name</label>
                  <input
                    value={form.name}
                    onChange={(e) =>
                      setForm({ ...form, name: e.target.value })
                    }
                    required
                  />
                </div>

                <div>
                  <label>Email</label>
                  <input
                    type="email"
                    value={form.email}
                    onChange={(e) =>
                      setForm({ ...form, email: e.target.value })
                    }
                    required
                  />
                </div>

                <div>
                  <label>Phone</label>
                  <input
                    value={form.phone}
                    onChange={(e) =>
                      setForm({ ...form, phone: e.target.value })
                    }
                    required
                  />
                </div>

                <div>
                  <label>License Number</label>
                  <input
                    value={form.license_number}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        license_number: e.target.value,
                      })
                    }
                    required
                  />
                </div>

                <div>
                  <label>Vehicle ID</label>
                  <input
                    type="number"
                    value={form.vehicle_id}
                    onChange={(e) =>
                      setForm({ ...form, vehicle_id: e.target.value })
                    }
                    placeholder="2"
                  />
                </div>

                <div>
                  <label>Status</label>
                  <select
                    value={form.status}
                    onChange={(e) =>
                      setForm({ ...form, status: e.target.value })
                    }
                  >
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="INACTIVE">INACTIVE</option>
                  </select>
                </div>
              </div>

              <button className="primary-button">
                Create Driver
              </button>
            </form>
          </div>
        )}

        <div className="table-card">
          <div className="table-header">
            <h2>Drivers</h2>
            <span>{drivers.length} drivers</span>
          </div>

          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Email</th>
                  <th>Phone</th>
                  <th>License</th>
                  <th>Vehicle</th>
                  <th>Status</th>
                  {isAdmin && <th>Action</th>}
                </tr>
              </thead>

              <tbody>
                {drivers.map((driver) => (
                  <tr key={driver.id}>
                    <td>#{driver.id}</td>
                    <td>{driver.name}</td>
                    <td>{driver.email}</td>
                    <td>{driver.phone}</td>
                    <td>{driver.license_number}</td>
                    <td>{driver.vehicle_id || "-"}</td>
                    <td>{driver.status}</td>

                    {isAdmin && (
                      <td>
                        <button
                          className="delete-button"
                          onClick={() => deleteDriver(driver.id)}
                        >
                          Delete
                        </button>
                      </td>
                    )}
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

export default Drivers;
