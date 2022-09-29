import React, { useState } from "react";
import { HiMenuAlt3 } from "react-icons/hi";
import { AiOutlinePlusCircle } from "react-icons/ai";
import { BsCardChecklist } from "react-icons/bs";
import { Link } from "react-router-dom";

const SideNavbar = () => {
    const menus = [
        {name:"Tickets",link: "/", icon: BsCardChecklist},
        {name:"Add Ticket",link:"/save", icon: AiOutlinePlusCircle},
    ];
    const [open, setOpen] = useState(true);

    return <div
        className={`bg-white min-h-screen ${
        open ? "w-72" : "w-16"
        } duration-500 text-gray-600 px-4`}
    >
        <div className="py-3 flex justify-end">
          <HiMenuAlt3
            size={26}
            className="cursor-pointer"
            onClick={() => setOpen(!open)}
          />
        </div>

        <div className="mt-4 flex flex-col gap-4 relative">
          {menus?.map((menu,i)=>(
            <Link
              to={menu?.link}
              key={i}
              className={` ${
                menu?.margin && "mt-5"
              } group flex items-center text-sm  gap-3.5 font-medium p-2 hover:bg-rose-400 hover:text-white rounded-md`}
            >
              <div>{React.createElement(menu?.icon, { size: "20" })}</div>
              <h2
                style={{
                  transitionDelay: `${i + 3}00ms`,
                }}
                className={`whitespace-pre duration-200 ${
                  !open && "opacity-0 translate-x-28 overflow-hidden"
                }`}
              >
                {menu?.name}
              </h2>
            </Link>
          ))}
        </div>
    </div>;    
}

export default SideNavbar