import { useEffect, useState } from "react";
import api from "../services/api";
import Layout from "../components/Layout";
import "../management.css";

function Vehicles() {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const user = JSON.parse(localStorage.getItem("user") || "{}");
  const isAdmin = user.role === "ADMIN";

  const [showForm, setShowForm] = useState(false);

  const [form, setForm] = useState({
    registration_number: "",
    make: "",
    model: "",
    vehicle_type: "",
    status: "ACTIVE",
  });

  const fetchVehicles = async () => {
    try {
      const token = localStorage.getItem("token");

      const response = await api.get("/vehicles", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      setVehicles(response.data);
    } catch (err) {
      setError(
        err.response?.data?.error || "Failed to load vehicles"
      );
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchVehicles();
  }, []);

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleCreate = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");

      await api.post("/vehicles", form, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      setForm({
        registration_number: "",
        make: "",
        model: "",
        vehicle_type: "",
        status: "ACTIVE",
      });

      setShowForm(false);
      fetchVehicles();
    } catch (err) {
      setError(
        err.response?.data?.error || "Failed to create vehicle"
      );
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Delete this vehicle?")) {
      return;
    }

    try {
      const token = localStorage.getItem("token");

      await api.delete(`/vehicles/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      fetchVehicles();
    } catch (err) {
      setError(
        err.response?.data?.error || "Failed to delete vehicle"
      );
    }
  };

  return (
    <Layout>
      <div className="management-page">
        <div className="page-header">
          <div>
            <h1>Vehicles</h1>
            <p>Manage your fleet vehicles</p>
          </div>

          {isAdmin && (
            <button
              className="primary-button"
              onClick={() => setShowForm(!showForm)}
            >
              {showForm ? "Cancel" : "+ Add Vehicle"}
            </button>
          )}
        </div>

        {error && (
          <div className="management-error">
            {error}
          </div>
        )}

        {showForm && isAdmin && (
          <div className="form-card">
            <h2>Add Vehicle</h2>

            <form onSubmit={handleCreate}>
              <div className="form-grid">
                <div>
                  <label>Registration Number</label>
                  <input
                    name="registration_number"
                    value={form.registration_number}
                    onChange={handleChange}
                    placeholder="GJ05AB1234"
                    required
                  />
                </div>

                <div>
                  <label>Make</label>
                  <input
                    name="make"
                    value={form.make}
                    onChange={handleChange}
                    placeholder="Tata"
                    required
                  />
                </div>

                <div>
                  <label>Model</label>
                  <input
                    name="model"
                    value={form.model}
                    onChange={handleChange}
                    placeholder="Nexon EV"
                    required
                  />
                </div>

                <div>
                  <label>Vehicle Type</label>
                  <select
                    name="vehicle_type"
                    value={form.vehicle_type}
                    onChange={handleChange}
                    required
                  >
                    <option value="">Select type</option>
                    <option value="SUV">SUV</option>
                    <option value="SEDAN">Sedan</option>
                    <option value="HATCHBACK">Hatchback</option>
                    <option value="TRUCK">Truck</option>
                    <option value="VAN">Van</option>
                  </select>
                </div>

                <div>
                  <label>Status</label>
                  <select
                    name="status"
                    value={form.status}
                    onChange={handleChange}
                  >
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="INACTIVE">INACTIVE</option>
                    <option value="MAINTENANCE">MAINTENANCE</option>
                  </select>
                </div>
              </div>

              <button
                type="submit"
                className="primary-button"
              >
                Create Vehicle
              </button>
            </form>
          </div>
        )}

        <div className="table-card">
          <div className="table-header">
            <h2>Fleet Vehicles</h2>
            <span>{vehicles.length} vehicles</span>
          </div>

          {loading ? (
            <div className="table-loading">
              Loading vehicles...
            </div>
          ) : vehicles.length === 0 ? (
            <div className="table-loading">
              No vehicles found.
            </div>
          ) : (
            <div className="table-container">
              <table>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Registration</th>
                    <th>Make</th>
                    <th>Model</th>
                    <th>Type</th>
                    <th>Status</th>
                    {isAdmin && <th>Action</th>}
                  </tr>
                </thead>

                <tbody>
                  {vehicles.map((vehicle) => (
                    <tr key={vehicle.id}>
                      <td>#{vehicle.id}</td>
                      <td>
                        <strong>
                          {vehicle.registration_number}
                        </strong>
                      </td>
                      <td>{vehicle.make}</td>
                      <td>{vehicle.model}</td>
                      <td>{vehicle.vehicle_type}</td>
                      <td>
                        <span
                          className={`status ${vehicle.status.toLowerCase()}`}
                        >
                          {vehicle.status}
                        </span>
                      </td>

                      {isAdmin && (
                        <td>
                          <button
                            className="delete-button"
                            onClick={() =>
                              handleDelete(vehicle.id)
                            }
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
          )}
        </div>
      </div>
    </Layout>
  );
}

export default Vehicles;
