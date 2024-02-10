import { configureStore } from '@reduxjs/toolkit'
import authSlice from './features/auth/authSlice'
import auctionSlice from './features/auction/auctionSlice'

export const store = configureStore({
    reducer: {
        auth: authSlice,
        auction: auctionSlice,
    },
})
