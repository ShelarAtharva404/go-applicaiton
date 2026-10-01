import { useEffect, useState } from "react";
import api from "../services/api";
import Layout from "../components/Layout";
import "../management.css";

function Trips() {
  const [trips, setTrips] = useState([]);
  const [error, setError] = useState("");
  const user = JSON.parse(localStorage.getItem("user") || "{}");
  const canCreate = user.role === "ADMIN" || user.role === "MANAGER";

  const [form, setForm] = useState({
    vehicle_id: "",
    driver_id: "",
    start_location: "",
    end_location: "",
    status: "PLANNED",
  });

  const loadTrips = async () => {
    try {
      const token = localStorage.getItem("token");

      const res = await api.get("/trips", {
        headers: { Authorization: `Bearer ${token}` },
      });

      setTrips(res.data || []);
    } catch (err) {
      setError(err.response?.data?.error || "Failed to load trips");
    }
  };

  useEffect(() => {
    loadTrips();
  }, []);

  const createTrip = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");

      await api.post(
        "/trips",
        {
          ...form,
          vehicle_id: Number(form.vehicle_id),
          driver_id: Number(form.driver_id),
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setForm({
        vehicle_id: "",
        driver_id: "",
        start_location: "",
        end_location: "",
        status: "PLANNED",
      });

      await loadTrips();
    } catch (err) {
      setError(err.response?.data?.error || "Failed to create trip");
    }
  };

  return (
    <Layout>
      <div className="management-page">
        <div className="page-header">
          <div>
            <h1>Trips</h1>
            <p>Manage fleet trips</p>
          </div>
        </div>

        {error && <div className="management-error">{error}</div>}

        {canCreate && (
          <div className="form-card">
            <h2>Create Trip</h2>

            <form onSubmit={createTrip}>
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
                    required
                  />
                </div>

                <div>
                  <label>Start Location</label>
                  <input
                    value={form.start_location}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        start_location: e.target.value,
                      })
                    }
                    required
                  />
                </div>

                <div>
                  <label>End Location</label>
                  <input
                    value={form.end_location}
                    onChange={(e) =>
                      setForm({
                        ...form,
                        end_location: e.target.value,
                      })
                    }
                    required
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
                    <option value="PLANNED">PLANNED</option>
                    <option value="IN_PROGRESS">IN_PROGRESS</option>
                    <option value="COMPLETED">COMPLETED</option>
                    <option value="CANCELLED">CANCELLED</option>
                  </select>
                </div>
              </div>

              <button className="primary-button">
                Create Trip
              </button>
            </form>
          </div>
        )}

        <div className="table-card">
          <div className="table-header">
            <h2>Trips</h2>
            <span>{trips.length} trips</span>
          </div>

          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Vehicle</th>
                  <th>Driver</th>
                  <th>From</th>
                  <th>To</th>
                  <th>Status</th>
                </tr>
              </thead>

              <tbody>
                {trips.map((trip) => (
                  <tr key={trip.id}>
                    <td>#{trip.id}</td>
                    <td>{trip.vehicle_id}</td>
                    <td>{trip.driver_id}</td>
                    <td>{trip.start_location}</td>
                    <td>{trip.end_location}</td>
                    <td>{trip.status}</td>
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

export default Trips;
