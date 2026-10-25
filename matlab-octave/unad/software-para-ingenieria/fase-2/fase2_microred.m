% =====================================================================
% RETO 2: MODELADO DEL BALANCE HIBRIDO
% SISTEMA INTELIGENTE DE GESTION ENERGETICA (SIGE)
% Curso: Software para Ingenieria (Codigo: 203036)
% AUTOR: Richard Adolfo Vega Almanza
% =====================================================================

clc; clear; close all;
hours = 1:24;

[solarEnergy, windEnergy] = generatePowerGenerationData();
totalEnergy = solarEnergy + windEnergy;

communityDemand = generateCommunityDemandData();

plotDay([
  PlotData(communityDemand, "r-*", 2, "Demanda de la comunidad"), ...
  PlotData(totalEnergy, "c-o", 2, "Generacion hibrida (solar y eolica)"), ...
]);

