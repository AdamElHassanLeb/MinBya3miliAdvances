import L from "leaflet";
import LocationIcon from '../assets/LocationIcon.png';


const MapIcon = new L.Icon ({
    iconUrl : LocationIcon,
    iconSize:     [70, 70], // size of the icon
    shadowSize:   [50, 64], // size of the shadow
    iconAnchor:   [30, 70], // point of the icon which will correspond to marker's location
    shadowAnchor: [4, 62],  // the same for the shadow
    popupAnchor:  [0, 0] // point from which the popup should open relative to the iconAnchorAnchor
})

export default MapIcon;