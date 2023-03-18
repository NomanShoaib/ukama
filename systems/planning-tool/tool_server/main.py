import uvicorn
from fastapi import FastAPI
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from utils.sites_coverage import SitesCoverage
from schema import SiteCoverage as SchemaSiteCoverage
from schema import Site as SchemaSite

from utils.sites_coverage import SitesCoverage
import os


app = FastAPI()

sitesCoverage = SitesCoverage()

@app.get("/")
async def root():
    return {"message": "Server is up and running."}

@app.post('/site_coverage/')
async def site_coverage(sites: list[SchemaSiteCoverage]):
    os.system("echo ran it")
    response = sitesCoverage.calculateCoverage(sites)
    json_compatible_item_data = jsonable_encoder(response)
    return JSONResponse(content=json_compatible_item_data) 

@app.post('/predict_links/')
async def predict_links(sitesCoverage: list[SchemaSiteCoverage]):
    json_compatible_item_data = jsonable_encoder(site)
    return JSONResponse(content=json_compatible_item_data)

@app.post('/get_elevation/')
async def get_elevation(sites: list[SchemaSite]):
    json_compatible_item_data = jsonable_encoder(sites)
    return JSONResponse(content=json_compatible_item_data)

# To run locally
if __name__ == '__main__':
    uvicorn.run(app, host='0.0.0.0', port=8000)