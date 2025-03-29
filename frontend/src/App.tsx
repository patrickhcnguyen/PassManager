import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Register from './components/Register/register'
import Login from './components/Login/login'
import Dashboard from './components/Dashboard/dashboard'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

function App() {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/dashboard" element={<Dashboard />} />
        </Routes>
      </Router>
    </QueryClientProvider>
  )
}

export default App
