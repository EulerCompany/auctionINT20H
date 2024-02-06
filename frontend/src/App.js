import { MainLayout } from './components/MainLayout.jsx'
import { Routes, Route } from 'react-router-dom'
import { ListingPage } from './pages/ListingPage.jsx'
import { LoginPage } from './pages/LoginPage.jsx'
import { RegistrationPage } from './pages/RegistrationPage.jsx'
import { useDispatch } from 'react-redux'
import { useEffect } from 'react'
import { getMe } from './redux/features/auth/authSlice.js'

function App() {
  const dispatch = useDispatch()

  useEffect(() => {
    // dispatch(getMe())
  }, [dispatch])

  return (
  <MainLayout>
    <Routes>
      <Route path='/' element={ <ListingPage /> } />
      <Route path='login' element={ <LoginPage /> } />
      <Route path='registration' element={ <RegistrationPage /> } />
    </Routes>
  </MainLayout>

)}

export default App;
