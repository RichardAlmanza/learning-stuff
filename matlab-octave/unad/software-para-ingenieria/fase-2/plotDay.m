% plotDay plots the PlotDataClass instances passed

function plotDay(dataList)
  hours = [1:24];

  hold on
  for i = 1:length(dataList)
    data = dataList(i);
    plot(hours, data.Data,...
        data.LineSpec, ...
        "DisplayName", data.Legend,...
        "LineWidth", data.LineWidth);
  end

  grid on;
  hold off;
  legend()

  % Add labels
  xlabel("Tiempo en horas");
  ylabel("Potencia en Kilovatios");

  % Add legend and title
  title("Balance Energetico de la Micro-Red SIGE");
  xticks(hours);
  xlim([1,24]);

  % End - Plot %
end
