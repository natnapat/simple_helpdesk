import React from "react";
import SideNavbar from "./SideNavbar";
import Ticket from "./Ticket";
import SaveTicket from "./SaveTicket";
import { BrowserRouter, Routes, Route } from "react-router-dom";

const App = () => {
    return (
        <div className="flex gap-6">
            <BrowserRouter>
                <SideNavbar />
                    <Routes>
                        <Route path="/" element={<Ticket />} />
                        <Route path="/save" element={<SaveTicket />}/>
                    </Routes>
            </BrowserRouter>
            
        </div>
    );
}

export default App

