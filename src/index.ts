import { separeHeaderAndBody } from './core/initial-preparations/initialPreparations';
import * as fs from "fs";
import path = require("path");

const EXAMPLE = true;

if(EXAMPLE){
    const agpmlExampleFile = fs.readFileSync(path.join(__dirname, "../assets/examples/agpml-example.agpml")) 
    separeHeaderAndBody(agpmlExampleFile.toString())
}


